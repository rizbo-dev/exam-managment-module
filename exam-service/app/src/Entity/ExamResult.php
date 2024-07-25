<?php

namespace App\Entity;

use App\Repository\ExamResultRepository;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: ExamResultRepository::class)]
class ExamResult
{
    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column]
    private ?int $id = null;

    #[ORM\OneToOne(cascade: ['persist', 'remove'])]
    private ?ExamStudent $examStudent = null;

    #[ORM\Column(length: 255)]
    private ?string $result = null;

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getExamStudent(): ?ExamStudent
    {
        return $this->examStudent;
    }

    public function setExamStudent(?ExamStudent $examStudent): static
    {
        $this->examStudent = $examStudent;

        return $this;
    }

    public function getResult(): ?string
    {
        return $this->result;
    }

    public function setResult(string $result): static
    {
        $this->result = $result;

        return $this;
    }
}
