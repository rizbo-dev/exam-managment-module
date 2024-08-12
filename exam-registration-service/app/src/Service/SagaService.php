<?php

namespace App\Service;

use App\Entity\ExamRegistration;
use App\Entity\SagaItem;
use App\Message\UserClassVerificationMessage;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\Messenger\MessageBusInterface;

class SagaService
{
    public function __construct(
        private readonly MessageBusInterface $messageBus,
        private readonly EntityManagerInterface $entityManager
    )
    {
    }

    public function dispatchUserClassVerificationSagaItem(SagaItem $sagaItem): void
    {
        $sagaItem
            ->setStartedAt(new \DateTimeImmutable())
            ->setStatus(SagaItem::IN_PROGRESS);

        $message = new UserClassVerificationMessage(
            studentId: $sagaItem->getExamRegistration()->getStudentId(),
            examId: $sagaItem->getExamRegistration()->getExamId(),
            examRegistrationId: $sagaItem->getExamRegistration()->getId()
        );

        $this->messageBus->dispatch($message);
        $this->entityManager->flush();
    }
}