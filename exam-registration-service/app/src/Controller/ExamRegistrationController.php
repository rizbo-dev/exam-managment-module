<?php

namespace App\Controller;

use App\Dto\CreateExamRegistrationDto;
use App\Service\ExamRegistrationService;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpKernel\Attribute\MapRequestPayload;
use Symfony\Component\Routing\Attribute\Route;

#[Route(path: "/exam-registrations", methods: "POST")]
class ExamRegistrationController extends AbstractController
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
        return $this->json('gg');
    }


}