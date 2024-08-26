<?php

namespace App\MessageHandler;

use App\Entity\Wallet;
use App\Message\UserUpdatedMessage;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;

#[AsMessageHandler]
class UserUpdateMessageHandler
{
    public function __construct(
        private readonly EntityManagerInterface $entityManager
    )
    {
    }

    public function __invoke(UserUpdatedMessage $message): void
    {
        $wallet = $this->entityManager->getRepository(Wallet::class)->findOneBy([
            'studentId' => $message->getUserId()
        ]);

        if ($wallet) {
            return;
        }

        $wallet = new Wallet();
        $wallet->setStudentId($message->getUserId())
            ->setBalance(0);

        $this->entityManager->persist($wallet);
        $this->entityManager->flush();
    }
}