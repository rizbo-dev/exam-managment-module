<?php

namespace App\Service\ResponseSagaItemResolver;

use App\Entity\ExamRegistration;
use App\Entity\SagaItem;
use App\Message\ResponseSagaItemMessage;
use App\Service\SagaItemDispatcher\SagaItemDispatcherService;
use App\Service\SagaItemService;
use Doctrine\ORM\EntityManagerInterface;
use Psr\Log\LoggerInterface;

class UserWalletValidationResponseSagaItemResolver implements ResponseSagaItemResolverInterface
{
    public function __construct(
        private readonly LoggerInterface $logger,
        private readonly EntityManagerInterface $entityManager,
        private readonly SagaItemDispatcherService $sagaItemDispatcherService
    )
    {
    }

    public function resolve(ResponseSagaItemMessage $responseSagaItemMessage): void
    {
        $payload = $responseSagaItemMessage->getPayload();
        $examRegistration = $this->entityManager->getRepository(ExamRegistration::class)->find($responseSagaItemMessage->getExamRegistrationId());
        $userClassVerificationSagaItem = SagaItemService::getUserWalletValidationSagaItem($examRegistration);


        if ($payload['isValid']) {
            $userClassVerificationSagaItem
                ->setStatus(SagaItem::FINISHED)
                ->setFinishedAt(new \DateTimeImmutable())
                ->setReturnedPayload($payload);

            $this->entityManager->flush();

            $nextSaga = SagaItemService::getNextItemForExecution($userClassVerificationSagaItem->getExamRegistration()->getSagaItems());

            $this->sagaItemDispatcherService->dispatch($nextSaga);
        }
    }

    public static function getResolverType(): string
    {
        return SagaItem::USER_WALLET_VALIDATION_SAGA_ITEM_TYPE;
    }
}