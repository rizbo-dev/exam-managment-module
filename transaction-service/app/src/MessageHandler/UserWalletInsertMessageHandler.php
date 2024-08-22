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

        $message = new ResponseSagaItemMessage(
            status: 'finished',
            sagaType: ResponseSagaItemMessage::USER_WALLET_INSERT_SAGA_ITEM_TYPE,
            examRegistrationId: $message->getExamRegistrationId(),
            payload: [
                'transactionId' => 1
            ]
        );

        $this->messageBus->dispatch($message);
    }
}