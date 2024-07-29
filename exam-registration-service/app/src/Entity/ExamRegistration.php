<?php

namespace App\Entity;

use App\Repository\ExamRegistrationRepository;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: ExamRegistrationRepository::class)]
class ExamRegistration
{
    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column]
    private ?int $id = null;

    #[ORM\Column]
    private ?int $studentId = null;

    #[ORM\Column]
    private ?int $examId = null;

    #[ORM\Column(length: 255)]
    private ?string $status = null;

    #[ORM\OneToOne(inversedBy: 'examRegistration', cascade: ['persist', 'remove'])]
    private ?ExamRegistrationSaga $sagaId = null;

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getStudentId(): ?int
    {
        return $this->studentId;
    }

    public function setStudentId(int $studentId): static
    {
        $this->studentId = $studentId;

        return $this;
    }

    public function getExamId(): ?int
    {
        return $this->examId;
    }

    public function setExamId(int $examId): static
    {
        $this->examId = $examId;

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

    public function getSagaId(): ?ExamRegistrationSaga
    {
        return $this->sagaId;
    }

    public function setSagaId(?ExamRegistrationSaga $sagaId): static
    {
        $this->sagaId = $sagaId;

        return $this;
    }
}
