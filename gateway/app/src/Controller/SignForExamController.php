<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\Routing\Attribute\Route;

#[Route('/sign-for-exam', name: 'app_sign_for_exam')]

class SignForExamController extends AbstractController
{
    public function __invoke()
    {

    }
}
