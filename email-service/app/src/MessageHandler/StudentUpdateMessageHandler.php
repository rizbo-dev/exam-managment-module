<?php

namespace App\MessageHandler;

use App\Message\StudentUpdateMessage;
use Symfony\Component\Mailer\MailerInterface;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;
use Symfony\Component\Mime\Email;

#[AsMessageHandler]
class StudentUpdateMessageHandler
{
    public function __construct(
        private readonly MailerInterface $mailer
    )
    {
    }

    public function __invoke(StudentUpdateMessage $message)
    {
        $email = (new Email())
            ->from('boris@boris.com')
            ->to('boris@boris.com')
            ->subject('Test')
            ->text('sending emails')
            ->html('<h1>Hello</h1>');

        $this->mailer->send($email);
    }
}