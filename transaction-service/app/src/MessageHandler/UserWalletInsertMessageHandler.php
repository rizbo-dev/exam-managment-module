<?php

namespace App\MessageHandler;

use App\Message\ResponseSagaItemMessage;
use App\Message\UserWalletInsertMessage;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;
use Symfony\Component\Messenger\MessageBusInterface;

#[AsMessageHandler]
readonly class UserWalletInsertMessageHandler
{
    public function __construct(
        private MessageBusInterface $messageBus
    )
    {
    }

    public function __invoke(UserWalletInsertMessage $message): void
    {
        /*
         * TODO:
         * Descrease amount on user wallet
         */

        $message = new ResponseSagaItemMessage();
        $message->setStatus('finished')
                ->setSagaType(ResponseSagaItemMessage::USER_WALLET_INSERT_SAGA_ITEM_TYPE)
                ->setExamRegistrationId($message->getExamRegistrationId())
                ->setPayload([
                    'transactionId' => 1
                ]);

        $this->messageBus->dispatch($message);
    }
}