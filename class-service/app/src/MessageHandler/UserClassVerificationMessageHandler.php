<?php

namespace App\MessageHandler;

use App\Entity\Student;
use App\Entity\StudyProgram;
use App\Message\ResponseSagaItemMessage;
use App\Message\UserClassVerificationMessage;
use Doctrine\ORM\EntityManagerInterface;
use Psr\Log\LoggerInterface;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;
use Symfony\Component\Messenger\MessageBusInterface;

#[AsMessageHandler]
readonly class UserClassVerificationMessageHandler
{
    public function __construct(
        private LoggerInterface $logger,
        private MessageBusInterface $messageBus,
        private EntityManagerInterface $entityManager
    )
    {
    }

    public function __invoke(UserClassVerificationMessage $classVerificationMessage)
    {
        $message = new ResponseSagaItemMessage();
        $student = $this->entityManager->getRepository(Student::class)->find($classVerificationMessage->getStudentId());

        if (!$student) {
            $message
                ->setSagaType(ResponseSagaItemMessage::USER_CLASS_VERIFICATION_SAGA_ITEM_TYPE)
                ->setStatus('finished')
                ->setExamRegistrationId($classVerificationMessage->getExamRegistrationId())
                ->setPayload([
                    'isValid' => false,
                    'message' => 'Student not found!'
                ]);
            return;
        }

        $found = false;

        /** @var StudyProgram $item */
        foreach ($student->getStudyProgram() as $item) {
            foreach ($item->getCourses() as $course) {
                if ($course->getId() === $classVerificationMessage->getCourseId()) {
                    $found = true;
                }
            }
        }

        if (!$found) {
            $message
                ->setSagaType(ResponseSagaItemMessage::USER_CLASS_VERIFICATION_SAGA_ITEM_TYPE)
                ->setStatus('finished')
                ->setExamRegistrationId($classVerificationMessage->getExamRegistrationId())
                ->setPayload([
                    'isValid' => false,
                    'message' => 'Student is not assigned to course!'
                ]);
            return;
        }

        $message
            ->setSagaType(ResponseSagaItemMessage::USER_CLASS_VERIFICATION_SAGA_ITEM_TYPE)
            ->setStatus('finished')
            ->setExamRegistrationId($classVerificationMessage->getExamRegistrationId())
            ->setPayload([
                'isValid' => true
            ]);

        $this->messageBus->dispatch($message);
    }
}