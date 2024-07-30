<?php

namespace App\Controller;

use App\Dto\CreateExamRegistrationDto;
use Symfony\Component\HttpKernel\Attribute\MapRequestPayload;
use Symfony\Component\Routing\Attribute\Route;

#[Route(path: "/exam-registrations", methods: "POST")]
class ExamRegistrationController
{
    public function __invoke(
        #[MapRequestPayload] CreateExamRegistrationDto $createExamRegistrationDto
    )
    {

    }


}