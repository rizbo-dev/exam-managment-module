<?php

namespace App\Service\ResponseSagaItemResolver;

use App\Entity\SagaItem;
use App\Message\ResponseSagaItemMessage;
use Psr\Log\LoggerInterface;

class UserWalletValidationResponseSagaItemResolver implements ResponseSagaItemResolverInterface
{
    public function __construct(
        private readonly LoggerInterface $logger
    )
    {
    }

    public function resolve(ResponseSagaItemMessage $responseSagaItemMessage): void
    {
        $this->logger->info('Handled in ' . SagaItem::USER_WALLET_VALIDATION_SAGA_ITEM_TYPE);
    }

    public static function getResolverType(): string
    {
        return SagaItem::USER_WALLET_VALIDATION_SAGA_ITEM_TYPE;
    }
}