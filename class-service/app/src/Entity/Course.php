<?php

namespace App\Entity;

use App\Repository\CourseRepository;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: CourseRepository::class)]
class Course
{
    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column]
    private ?int $id = null;

    #[ORM\Column]
    private ?int $teacherId = null;

    #[ORM\Column(length: 255)]
    private ?string $name = null;

    /**
     * @var Collection<int, StudyProgram>
     */
    #[ORM\ManyToMany(targetEntity: StudyProgram::class, inversedBy: 'courses')]
    private Collection $studyProgram;

    public function __construct()
    {
        $this->studyProgram = new ArrayCollection();
    }

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getTeacherId(): ?int
    {
        return $this->teacherId;
    }

    public function setTeacherId(int $teacherId): static
    {
        $this->teacherId = $teacherId;

        return $this;
    }

    public function getName(): ?string
    {
        return $this->name;
    }

    public function setName(string $name): static
    {
        $this->name = $name;

        return $this;
    }

    /**
     * @return Collection<int, StudyProgram>
     */
    public function getStudyProgram(): Collection
    {
        return $this->studyProgram;
    }

    public function addStudyProgram(StudyProgram $studyProgram): static
    {
        if (!$this->studyProgram->contains($studyProgram)) {
            $this->studyProgram->add($studyProgram);
        }

        return $this;
    }

    public function removeStudyProgram(StudyProgram $studyProgram): static
    {
        $this->studyProgram->removeElement($studyProgram);

        return $this;
    }
}