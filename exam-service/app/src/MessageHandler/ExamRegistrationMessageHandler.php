<?php

namespace App\MessageHandler;

use App\Message\ExamRegistrationMessage;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;

#[AsMessageHandler]
class ExamRegistrationMessageHandler
{
    public function __invoke(ExamRegistrationMessage $message)
    {
        // TODO: Implement __invoke() method.
    }
}