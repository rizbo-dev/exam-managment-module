<?php

namespace App\Entity;

use App\Repository\SagaItemRepository;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: SagaItemRepository::class)]
class SagaItem
{
    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column]
    private ?int $id = null;

    #[ORM\Column(length: 255)]
    private ?string $type = null;

    #[ORM\Column]
    private ?\DateTimeImmutable $startedAt = null;

    #[ORM\Column(nullable: true)]
    private ?\DateTimeImmutable $finishedAt = null;

    #[ORM\Column(length: 255)]
    private ?string $status = null;

    #[ORM\OneToOne(mappedBy: 'userClassVerificationSagaItem', cascade: ['persist', 'remove'])]
    private ?ExamRegistrationSaga $examRegistrationSaga = null;

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getType(): ?string
    {
        return $this->type;
    }

    public function setType(string $type): static
    {
        $this->type = $type;

        return $this;
    }

    public function getStartedAt(): ?\DateTimeImmutable
    {
        return $this->startedAt;
    }

    public function setStartedAt(\DateTimeImmutable $startedAt): static
    {
        $this->startedAt = $startedAt;

        return $this;
    }

    public function getFinishedAt(): ?\DateTimeImmutable
    {
        return $this->finishedAt;
    }

    public function setFinishedAt(?\DateTimeImmutable $finishedAt): static
    {
        $this->finishedAt = $finishedAt;

        return $this;
    }

    public function getStatus(): ?string
    {
        return $this->status;
    }

    public function setStatus(string $status): static
    {
        $this->status = $status;

        return $this;
    }

    public function getExamRegistrationSaga(): ?ExamRegistrationSaga
    {
        return $this->examRegistrationSaga;
    }

    public function setExamRegistrationSaga(ExamRegistrationSaga $examRegistrationSaga): static
    {
        // set the owning side of the relation if necessary
        if ($examRegistrationSaga->getUserClassVerificationSagaItem() !== $this) {
            $examRegistrationSaga->setUserClassVerificationSagaItem($this);
        }

        $this->examRegistrationSaga = $examRegistrationSaga;

        return $this;
    }
}
