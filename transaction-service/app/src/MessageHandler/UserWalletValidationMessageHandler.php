<?php

namespace App\MessageHandler;

use App\Message\ResponseSagaItemMessage;
use App\Message\UserWalletValidationMessage;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;
use Symfony\Component\Messenger\MessageBusInterface;

#[AsMessageHandler]
readonly class UserWalletValidationMessageHandler
{
    public function __construct(
        private MessageBusInterface $messageBus,

    )
    {
    }

    public function __invoke(UserWalletValidationMessage $message)
    {
        /**
         * TODO:
         * 1. Fetch exam cost
         * 2. Compare it to wallet
         * 3. If there is not enough credit in wallet trigger isValid false
         * 4. If there is enough credit in wallet trigger isValid true
         */

        $message = new ResponseSagaItemMessage(
            status: 'finished',
            sagaType: ResponseSagaItemMessage::USER_WALLET_VALIDATION_SAGA_ITEM_TYPE,
            examRegistrationId: $message->getExamRegistrationId(),
            payload: [
                'isValid' => true,
                'examCreditValue' => 1000
            ]
        );

        $this->messageBus->dispatch($message);
    }
}