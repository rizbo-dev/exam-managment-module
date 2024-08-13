<?php

namespace App\Service\SagaItemDispatcher;

use App\Entity\SagaItem;

class UserWalletValidationSagaItemDispatcher implements SagaItemDispatcherInterface
{

    public function dispatch(SagaItem $sagaItem): void
    {
        // TODO: Implement dispatch() method.
    }

    public static function getDispatcherType(): string
    {
        return SagaItem::USER_WALLET_VALIDATION_SAGA_ITEM_TYPE;
    }
}