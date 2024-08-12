<?php

namespace App\Service;

use App\Dto\CreateExamRegistrationDto;
use App\Entity\ExamRegistration;
use App\Entity\SagaItem;
use Doctrine\ORM\EntityManagerInterface;
use http\Exception\RuntimeException;

readonly class ExamRegistrationService
{
    public function __construct(
        private EntityManagerInterface $entityManager,
        private SagaService $sagaService
    )
    {
    }

    public function initExamRegistrationSaga(CreateExamRegistrationDto $createExamRegistrationDto): ExamRegistration
    {
        $examRegistration = new ExamRegistration();

        $examRegistration
            ->setStudentId($createExamRegistrationDto->getStudentId())
            ->setExamId($createExamRegistrationDto->getExamId());

        $this->entityManager->persist($examRegistration);
        $this->initSagaItemsForExamRegistration($examRegistration);

        $this->entityManager->flush();

        $userClassVerificationSagaItem = SagaItemService::getUserClassVerificationSagaItem($examRegistration);
        if (!$userClassVerificationSagaItem) {
            throw new RuntimeException('handle exception');
        }
        $this->sagaService->dispatchUserClassVerificationSagaItem($userClassVerificationSagaItem);

        return $examRegistration;
    }

    public function initSagaItemsForExamRegistration(ExamRegistration $examRegistration): void
    {
        foreach (SagaItem::ITEMS as $item) {
            $sagaItem = new SagaItem();
            $sagaItem->setSagaType($item['sagaType'])
                     ->setExecutionOrder($item['executionOrder'])
                    ->setType($item['type']);

            $this->entityManager->persist($sagaItem);
            $examRegistration->addSagaItem($sagaItem);
        }
    }
}