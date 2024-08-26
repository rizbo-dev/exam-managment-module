<?php

namespace App\MessageHandler;

use App\Entity\Transaction;
use App\Entity\Wallet;
use App\Message\ResponseSagaItemMessage;
use App\Message\UserWalletInsertMessage;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;
use Symfony\Component\Messenger\MessageBusInterface;

#[AsMessageHandler]
readonly class UserWalletInsertMessageHandler
{
    public function __construct(
        private MessageBusInterface $messageBus,
        private EntityManagerInterface $entityManager
    )
    {
    }

    public function __invoke(UserWalletInsertMessage $message): void
    {
        /** @var Wallet $userWallet */
        $userWallet = $this->entityManager->getRepository(Wallet::class)->findOneBy([
            'studentId' => $message->getStudentId()
        ]);

        $transaction = new Transaction();
        $transaction
            ->setStatus('processed')
            ->setAmount($message->getAmount())
            ->setType('outcome');

        $this->entityManager->persist($transaction);

        $userWallet->addTransaction($transaction);

        $this->entityManager->flush();

        $message = new ResponseSagaItemMessage(
            status: 'finished',
            sagaType: ResponseSagaItemMessage::USER_WALLET_INSERT_SAGA_ITEM_TYPE,
            examRegistrationId: $message->getExamRegistrationId(),
            payload: [
                'transactionId' => $transaction->getId()
            ]
        );

        $this->messageBus->dispatch($message);
    }
}