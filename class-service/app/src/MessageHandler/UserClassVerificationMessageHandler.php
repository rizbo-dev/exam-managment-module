<?php

namespace App\MessageHandler;

use App\Message\ResponseSagaItemMessage;
use App\Message\UserClassVerificationMessage;
use Psr\Log\LoggerInterface;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;
use Symfony\Component\Messenger\MessageBusInterface;

#[AsMessageHandler]
readonly class UserClassVerificationMessageHandler
{
    public function __construct(
        private LoggerInterface $logger,
        private MessageBusInterface $messageBus
    )
    {
    }

    public function __invoke(UserClassVerificationMessage $classVerificationMessage)
    {
        $this->logger->info('Consumed');

        //TODO handle this action

        $message = new ResponseSagaItemMessage();
        $message
            ->setSagaType(ResponseSagaItemMessage::USER_CLASS_VERIFICATION_SAGA_ITEM_TYPE)
            ->setStatus('finished')
            ->setExamRegistrationId($classVerificationMessage->getExamRegistrationId())
            ->setPayload([
                'isValid' => true
            ]);

        $this->messageBus->dispatch($message);
    }
}