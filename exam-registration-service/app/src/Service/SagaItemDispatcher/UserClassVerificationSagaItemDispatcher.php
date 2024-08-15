<?php

namespace App\Service\SagaItemDispatcher;

use App\Entity\SagaItem;
use App\Message\UserClassVerificationMessage;
use Symfony\Component\Messenger\MessageBusInterface;

readonly class UserClassVerificationSagaItemDispatcher implements SagaItemDispatcherInterface
{
    public function __construct(
        private MessageBusInterface $messageBus
    )
    {
    }


    public function dispatch(SagaItem $sagaItem): void
    {
        $message = new UserClassVerificationMessage(
            studentId: $sagaItem->getExamRegistration()->getStudentId(),
            examId: $sagaItem->getExamRegistration()->getExamId(),
            examRegistrationId: $sagaItem->getExamRegistration()->getId()
        );

        $this->messageBus->dispatch($message);
    }

    public static function getDispatcherType(): string
    {
        return SagaItem::USER_CLASS_VERIFICATION_SAGA_ITEM_TYPE;
    }
}