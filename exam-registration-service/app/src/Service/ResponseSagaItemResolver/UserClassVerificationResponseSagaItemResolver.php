<?php

namespace App\Service\ResponseSagaItemResolver;

use App\Entity\ExamRegistration;
use App\Entity\SagaItem;
use App\Message\ResponseSagaItemMessage;
use App\Service\SagaItemService;
use Doctrine\ORM\EntityManagerInterface;
use Psr\Log\LoggerInterface;

readonly class UserClassVerificationResponseSagaItemResolver implements ResponseSagaItemResolverInterface
{
    public function __construct(
        private LoggerInterface $logger,
        private EntityManagerInterface $entityManager
    )
    {
    }

    public function resolve(ResponseSagaItemMessage $responseSagaItemMessage): void
    {
        $payload = $responseSagaItemMessage->getPayload();

        $examRegistration = $this->entityManager->getRepository(ExamRegistration::class)->find($responseSagaItemMessage->getExamRegistrationId());

        $userClassVerificationSagaItem = SagaItemService::getUserClassVerificationSagaItem($examRegistration);


        if ($payload['isValid']) {
            $userClassVerificationSagaItem
                ->setStatus(SagaItem::FINISHED)
                ->setFinishedAt(new \DateTimeImmutable());

            $this->entityManager->flush();

            $nextSaga = SagaItemService::getNextItemForExecution($userClassVerificationSagaItem);
        }
    }

    public static function getResolverType(): string
    {
        return SagaItem::USER_CLASS_VERIFICATION_SAGA_ITEM_TYPE;
    }
}