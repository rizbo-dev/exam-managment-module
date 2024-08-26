<?php

namespace App\MessageHandler;

use App\Entity\Exam;
use App\Entity\ExamStudent;
use App\Entity\SagaItem;
use App\Message\ExamRegistrationMessage;
use App\Message\ResponseSagaItemMessage;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;
use Symfony\Component\Messenger\MessageBusInterface;

#[AsMessageHandler]
readonly class ExamRegistrationMessageHandler
{
    public function __construct(
        private MessageBusInterface    $messageBus,
        private EntityManagerInterface $entityManager
    )
    {
    }

    public function __invoke(ExamRegistrationMessage $message): void
    {
        try {
            $examStudent = new ExamStudent();
            $exam = $this->entityManager->getRepository(Exam::class)->find($message->getExamId());

            $examStudent->setExam($exam)
                ->setStudentId($message->getStudentId());

            $this->entityManager->persist($examStudent);
            $this->entityManager->flush();

            $responseSagaItem = new ResponseSagaItemMessage();
            $responseSagaItem
                ->setExamRegistrationId($message->getExamRegistrationId())
                ->setPayload([
                    'examStudentId' => $examStudent->getId()
                ])
                ->setStatus('finished')
                ->setSagaType('examRegistrationSagaItem');

            $this->messageBus->dispatch($responseSagaItem);
        } catch (\Throwable $throwable) {
            $responseSagaItem = new ResponseSagaItemMessage();
            $responseSagaItem
                ->setExamRegistrationId($message->getExamRegistrationId())
                ->setPayload([
                    'message' => $throwable->getMessage()
                ])
                ->setStatus('finished')
                ->setSagaType('examRegistrationSagaItem');
            $this->messageBus->dispatch($responseSagaItem);
        }
    }
}