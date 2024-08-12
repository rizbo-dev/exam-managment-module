<?php

namespace App\MessageHandler;

use App\Message\ResponseSagaItemMessage;
use Psr\Log\LoggerInterface;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;

#[AsMessageHandler]
readonly class ResponseSagaItemMessageHandler
{
    public function __construct(
        private LoggerInterface $logger
    )
    {
    }

    public function __invoke(ResponseSagaItemMessage $responseSagaItemMessage): void
    {
        $this->logger->info('Consumed');
    }
}