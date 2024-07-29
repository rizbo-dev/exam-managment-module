<?php

namespace App\Entity;

use App\Repository\ExamRegistrationSagaRepository;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: ExamRegistrationSagaRepository::class)]
class ExamRegistrationSaga
{
    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column]
    private ?int $id = null;

    #[ORM\OneToOne(mappedBy: 'sagaId', cascade: ['persist', 'remove'])]
    private ?ExamRegistration $examRegistration = null;

    #[ORM\OneToOne(inversedBy: 'examRegistrationSaga', cascade: ['persist', 'remove'])]
    #[ORM\JoinColumn(nullable: false)]
    private ?SagaItem $userClassVerificationSagaItem = null;

    #[ORM\OneToOne(inversedBy: 'examRegistrationSaga', cascade: ['persist', 'remove'])]
    #[ORM\JoinColumn(nullable: false)]
    private ?SagaItem $userWalletValidationSagaItem = null;

    #[ORM\OneToOne(inversedBy: 'examRegistrationSaga', cascade: ['persist', 'remove'])]
    #[ORM\JoinColumn(nullable: false)]
    private ?SagaItem $examValidationSagaItem = null;

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getExamRegistration(): ?ExamRegistration
    {
        return $this->examRegistration;
    }

    public function setExamRegistration(?ExamRegistration $examRegistration): static
    {
        // unset the owning side of the relation if necessary
        if ($examRegistration === null && $this->examRegistration !== null) {
            $this->examRegistration->setSagaId(null);
        }

        // set the owning side of the relation if necessary
        if ($examRegistration !== null && $examRegistration->getSagaId() !== $this) {
            $examRegistration->setSagaId($this);
        }

        $this->examRegistration = $examRegistration;

        return $this;
    }

    public function getUserClassVerificationSagaItem(): ?SagaItem
    {
        return $this->userClassVerificationSagaItem;
    }

    public function setUserClassVerificationSagaItem(SagaItem $userClassVerificationSagaItem): static
    {
        $this->userClassVerificationSagaItem = $userClassVerificationSagaItem;

        return $this;
    }

    public function getUserWalletValidationSagaItem(): ?SagaItem
    {
        return $this->userWalletValidationSagaItem;
    }

    public function setUserWalletValidationSagaItem(SagaItem $userWalletValidationSagaItem): static
    {
        $this->userWalletValidationSagaItem = $userWalletValidationSagaItem;

        return $this;
    }

    public function getExamValidationSagaItem(): ?SagaItem
    {
        return $this->examValidationSagaItem;
    }

    public function setExamValidationSagaItem(SagaItem $examValidationSagaItem): static
    {
        $this->examValidationSagaItem = $examValidationSagaItem;

        return $this;
    }
}
