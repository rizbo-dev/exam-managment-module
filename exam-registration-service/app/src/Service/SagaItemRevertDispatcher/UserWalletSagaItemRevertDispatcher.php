<?php

namespace App\Service\SagaItemRevertDispatcher;

use App\Entity\SagaItem;
use App\Message\UserWalletRevertMessage;
use Symfony\Component\Messenger\MessageBusInterface;

class UserWalletSagaItemRevertDispatcher implements SagaItemRevertDispatcherInterface
{
    public function __construct(
        private readonly MessageBusInterface $messageBus
    )
    {
    }

    public function revert(SagaItem $sagaItem): void
    {
        $userWalletRevertMessage = new UserWalletRevertMessage();
        $userWalletRevertMessage->setTransactionId($sagaItem->getReturnedPayload()['transactionId']);
        $this->messageBus->dispatch($userWalletRevertMessage);
    }

    public static function getDispatcherType(): string
    {
        return SagaItem::USER_WALLET_INSERT_SAGA_ITEM_TYPE;
    }
}