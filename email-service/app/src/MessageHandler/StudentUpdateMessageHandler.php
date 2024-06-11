<?php

namespace App\MessageHandler;

use App\Message\StudentUpdateMessage;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;

#[AsMessageHandler]
class StudentUpdateMessageHandler
{
    public function __invoke(StudentUpdateMessage $message)
    {
        dump($message);
    }
}