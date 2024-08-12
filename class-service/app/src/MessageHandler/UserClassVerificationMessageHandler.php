<?php

namespace App\MessageHandler;

use App\Message\UserClassVerificationMessage;
use Psr\Log\LoggerInterface;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;

#[AsMessageHandler]
class UserClassVerificationMessageHandler
{
    public function __construct(
        private  readonly LoggerInterface $logger
    )
    {
    }

    public function __invoke(UserClassVerificationMessage $classVerificationMessage)
    {
        $this->logger->info('Consumed');
    }
}