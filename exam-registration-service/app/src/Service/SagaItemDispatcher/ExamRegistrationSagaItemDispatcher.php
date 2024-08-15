<?php

namespace App\Service\SagaItemDispatcher;

use App\Entity\SagaItem;
use App\Message\ExamRegistrationMessage;
use Symfony\Component\Messenger\MessageBusInterface;

readonly class ExamRegistrationSagaItemDispatcher implements SagaItemDispatcherInterface
{
    public function __construct(
        private MessageBusInterface $messageBus
    )
    {

    }


    public function dispatch(SagaItem $sagaItem): void
    {
        $message = new ExamRegistrationMessage(
            studentId: $sagaItem->getExamRegistration()->getStudentId(),
            examId: $sagaItem->getExamRegistration()->getExamId(),
            examRegistrationId: $sagaItem->getExamRegistration()->getId()
        );

        $this->messageBus->dispatch($message);
    }

    public static function getDispatcherType(): string
    {
        return SagaItem::EXAM_REGISTRATION_SAGA_ITEM_TYPE;
    }
}