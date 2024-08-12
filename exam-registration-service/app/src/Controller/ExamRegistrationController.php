<?php

namespace App\Controller;

use App\Dto\CreateExamRegistrationDto;
use App\Service\ExamRegistrationService;
use Symfony\Component\HttpKernel\Attribute\MapRequestPayload;
use Symfony\Component\Routing\Attribute\Route;

#[Route(path: "/exam-registrations", methods: "POST")]
class ExamRegistrationController
{
    public function __construct(
        private readonly ExamRegistrationService $examRegistrationService
    )
    {
    }

    public function __invoke(
        #[MapRequestPayload] CreateExamRegistrationDto $createExamRegistrationDto
    )
    {
        $createdSaga = $this->examRegistrationService->initExamRegistrationSaga($createExamRegistrationDto);
    }


}