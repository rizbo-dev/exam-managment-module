<?php

namespace App\MessageHandler;

use App\Entity\SagaItem;
use App\Message\ExamRegistrationMessage;
use App\Message\ResponseSagaItemMessage;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;
use Symfony\Component\Messenger\MessageBusInterface;

#[AsMessageHandler]
class ExamRegistrationMessageHandler
{
    public function __construct(
        private readonly MessageBusInterface $messageBus
    )
    {
    }

    public function __invoke(ExamRegistrationMessage $message)
    {
        $responseSagaItem = new ResponseSagaItemMessage();
        $responseSagaItem
            ->setExamRegistrationId($message->getExamRegistrationId())
            ->setPayload([
                'examStudentId' => 12
            ])
            ->setStatus('finished')
            ->setSagaType('examRegistrationSagaItem');

        $this->messageBus->dispatch($responseSagaItem);
    }
}