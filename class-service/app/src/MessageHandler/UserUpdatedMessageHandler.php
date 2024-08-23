<?php

namespace App\MessageHandler;

use App\Entity\Department;
use App\Entity\Student;
use App\Entity\StudyProgram;
use App\Message\UserUpdatedMessage;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;
use Symfony\Contracts\HttpClient\HttpClientInterface;

#[AsMessageHandler]
class UserUpdatedMessageHandler
{
    public function __construct(
        private readonly HttpClientInterface $userClient,
        private readonly EntityManagerInterface $entityManager
    )
    {
    }

    public function __invoke(UserUpdatedMessage $message)
    {
        $res = $this->userClient->request(Request::METHOD_GET, '/api/users/' . $message->getUserId())->toArray();

        if ($res['role'] !== 'student') {
            return;
        }

        $student = $this->entityManager->getRepository(Student::class)->find($message->getUserId());

        if (!$student) {
            $student = new Student();
            $this->entityManager->persist($student);
        }

        $department = $res['departmentId'] ? $this->entityManager->getRepository(Department::class)->find($res['departmentId']) : null;
        $studyProgram = $res['studyProgramId'] ? $this->entityManager->getRepository(StudyProgram::class)->find($res['studyProgramId']) : null;


        $student
            ->setId($res['id'])
            ->setFirstname($res['firstName'])
            ->setLastname($res['lastName'])
            ->setDepartment($department)
            ->setStudyProgram($studyProgram);

        $this->entityManager->flush();
    }
}