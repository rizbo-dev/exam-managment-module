<?php

namespace App\Service\ResponseSagaItemResolver;

use App\Entity\ExamRegistration;
use App\Entity\SagaItem;
use App\Message\ResponseSagaItemMessage;
use App\Service\SagaItemService;
use Doctrine\ORM\EntityManagerInterface;

class ExamRegistrationResponseSagaItemResolver implements ResponseSagaItemResolverInterface
{
    public function __construct(
        private readonly EntityManagerInterface $entityManager
    )
    {
    }

    public function resolve(ResponseSagaItemMessage $responseSagaItemMessage): void
    {
        $payload = $responseSagaItemMessage->getPayload();
        $examRegistration = $this->entityManager->getRepository(ExamRegistration::class)->find($responseSagaItemMessage->getExamRegistrationId());
        $examRegistrationSagaItem = SagaItemService::getExamRegistrationSagaItem($examRegistration);

        $examRegistrationSagaItem
            ->setStatus(SagaItem::FINISHED)
            ->setFinishedAt(new \DateTimeImmutable())
            ->setReturnedPayload($payload);

        $this->entityManager->flush();
    }

    public static function getResolverType(): string
    {
        return SagaItem::EXAM_REGISTRATION_SAGA_ITEM_TYPE;
    }
}