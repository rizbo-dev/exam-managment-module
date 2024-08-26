<?php

namespace App\Service\ResponseSagaItemResolver;

use App\Entity\ExamRegistration;
use App\Entity\SagaItem;
use App\Message\ResponseSagaItemMessage;
use App\Service\SagaItemRevertDispatcher\SagaItemRevertDispatcherService;
use App\Service\SagaItemService;
use Doctrine\ORM\EntityManagerInterface;

class ExamRegistrationResponseSagaItemResolver implements ResponseSagaItemResolverInterface
{
    public function __construct(
        private readonly EntityManagerInterface $entityManager,
        private readonly SagaItemRevertDispatcherService $dispatcherService
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

        if (isset($payload['message'])) {
            SagaItemService::markSagaItemsAsCanceled($examRegistration);
            foreach ($examRegistration->getSagaItems() as $sagaItem) {
                $this->dispatcherService->revert($sagaItem);
            }
        }
    }

    public static function getResolverType(): string
    {
        return SagaItem::EXAM_REGISTRATION_SAGA_ITEM_TYPE;
    }
}