<?php

namespace App\Service\ResponseSagaItemResolver;

use App\Entity\ExamRegistration;
use App\Entity\SagaItem;
use App\Message\ResponseSagaItemMessage;
use App\Service\SagaItemDispatcher\SagaItemDispatcherService;
use App\Service\SagaItemService;
use Doctrine\ORM\EntityManagerInterface;

readonly class UserWalletValidationResponseSagaItemResolver implements ResponseSagaItemResolverInterface
{
    public function __construct(
        private EntityManagerInterface    $entityManager,
        private SagaItemDispatcherService $sagaItemDispatcherService
    )
    {
    }

    public function resolve(ResponseSagaItemMessage $responseSagaItemMessage): void
    {
        $payload = $responseSagaItemMessage->getPayload();
        $examRegistration = $this->entityManager->getRepository(ExamRegistration::class)->find($responseSagaItemMessage->getExamRegistrationId());
        $userWalletValidationSagaItem = SagaItemService::getUserWalletValidationSagaItem($examRegistration);
        $userWalletValidationSagaItem
            ->setStatus(SagaItem::FINISHED)
            ->setFinishedAt(new \DateTimeImmutable())
            ->setReturnedPayload($payload);

        $this->entityManager->flush();

        if (!$payload['isValid']) {
            SagaItemService::markSagaItemsAsCanceled($examRegistration);
            $this->entityManager->flush();
            return;
        }

        $nextSaga = SagaItemService::getNextItemForExecution($userWalletValidationSagaItem->getExamRegistration()->getSagaItems());

        $this->sagaItemDispatcherService->dispatch($nextSaga);
    }
    public static function getResolverType(): string
    {
        return SagaItem::USER_WALLET_VALIDATION_SAGA_ITEM_TYPE;
    }

}