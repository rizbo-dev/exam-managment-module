<?php

namespace App\MessageHandler;

use App\Entity\Transaction;
use App\Message\UserWalletRevertMessage;
use App\Service\WalletService;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;

#[AsMessageHandler]
readonly class UserWalletRevertMessageHandler
{
    public function __construct(
        private EntityManagerInterface $entityManager
    )
    {
    }

    public function __invoke(
        UserWalletRevertMessage $message
    )
    {
        $transaction = $this->entityManager->getRepository(Transaction::class)->find($message->getTransactionId());
        $wallet = $transaction->getWallet();

        $this->entityManager->remove($transaction);
        $this->entityManager->flush();
        WalletService::syncWalletBalance($wallet);
        $this->entityManager->flush();
    }
}