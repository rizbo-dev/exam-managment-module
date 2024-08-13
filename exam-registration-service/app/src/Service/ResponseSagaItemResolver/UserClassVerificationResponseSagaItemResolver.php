<?php

namespace App\Service\ResponseSagaItemResolver;

use App\Entity\SagaItem;
use App\Message\ResponseSagaItemMessage;
use Psr\Log\LoggerInterface;

readonly class UserClassVerificationResponseSagaItemResolver implements ResponseSagaItemResolverInterface
{
    public function __construct(
        private LoggerInterface $logger
    )
    {
    }

    public function resolve(ResponseSagaItemMessage $responseSagaItemMessage): void
    {
        $this->logger->info('handling in ' . SagaItem::USER_CLASS_VERIFICATION_SAGA_ITEM_TYPE);
    }

    public static function getResolverType(): string
    {
        return SagaItem::USER_CLASS_VERIFICATION_SAGA_ITEM_TYPE;
    }
}