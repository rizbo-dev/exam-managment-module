<?php

namespace App\MessageHandler;

use App\Entity\Wallet;
use App\Message\ResponseSagaItemMessage;
use App\Message\UserWalletValidationMessage;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;
use Symfony\Component\Messenger\MessageBusInterface;
use Symfony\Contracts\HttpClient\HttpClientInterface;

#[AsMessageHandler]
readonly class UserWalletValidationMessageHandler
{
    public function __construct(
        private MessageBusInterface $messageBus,
        private HttpClientInterface $examClient,
        private EntityManagerInterface $entityManager
    )
    {
    }

    public function __invoke(UserWalletValidationMessage $message)
    {
        $exam = $this->examClient->request(Request::METHOD_GET, '/api/exams/' . $message->getExamId())->toArray();

        $examCost = $exam['cost'];

        /** @var Wallet|null $wallet */
        $wallet = $this->entityManager->getRepository(Wallet::class)->findOneBy([
            'studentId' => $message->getStudentId()
        ]);

        if (!$wallet) {
            $message = new ResponseSagaItemMessage(
                status: 'finished',
                sagaType: ResponseSagaItemMessage::USER_WALLET_VALIDATION_SAGA_ITEM_TYPE,
                examRegistrationId: $message->getExamRegistrationId(),
                payload: [
                    'isValid' => false,
                    'message' => 'Student wallet not found!'
                ]
            );

            $this->messageBus->dispatch($message);

            return;
        }

        if (($wallet->getBalance() - $examCost) < 0) {
            $message = new ResponseSagaItemMessage(
                status: 'finished',
                sagaType: ResponseSagaItemMessage::USER_WALLET_VALIDATION_SAGA_ITEM_TYPE,
                examRegistrationId: $message->getExamRegistrationId(),
                payload: [
                    'isValid' => false,
                    'message' => 'Not enough funds on wallet'
                ]
            );

            $this->messageBus->dispatch($message);

            return;
        }


        $message = new ResponseSagaItemMessage(
            status: 'finished',
            sagaType: ResponseSagaItemMessage::USER_WALLET_VALIDATION_SAGA_ITEM_TYPE,
            examRegistrationId: $message->getExamRegistrationId(),
            payload: [
                'isValid' => true,
                'examCreditValue' => $examCost
            ]
        );

        $this->messageBus->dispatch($message);
    }
}