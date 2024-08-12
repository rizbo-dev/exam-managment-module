<?php

namespace App\Entity;

use App\Repository\DepartmentRepository;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: DepartmentRepository::class)]
class Department
{
    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column]
    private ?int $id = null;

    #[ORM\Column(length: 255)]
    private ?string $name = null;

    /**
     * @var Collection<int, StudyProgram>
     */
    #[ORM\OneToMany(targetEntity: StudyProgram::class, mappedBy: 'department')]
    private Collection $studyPrograms;

    /**
     * @var Collection<int, Student>
     */
    #[ORM\OneToMany(targetEntity: Student::class, mappedBy: 'department')]
    private Collection $students;

    public function __construct()
    {
        $this->studyPrograms = new ArrayCollection();
        $this->students = new ArrayCollection();
    }

    public function getId(): ?int
    {
        return $this->id;
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
    public function getStudyPrograms(): Collection
    {
        return $this->studyPrograms;
    }

    public function addStudyProgram(StudyProgram $studyProgram): static
    {
        if (!$this->studyPrograms->contains($studyProgram)) {
            $this->studyPrograms->add($studyProgram);
            $studyProgram->setDepartment($this);
        }

        return $this;
    }

    public function removeStudyProgram(StudyProgram $studyProgram): static
    {
        if ($this->studyPrograms->removeElement($studyProgram)) {
            // set the owning side to null (unless already changed)
            if ($studyProgram->getDepartment() === $this) {
                $studyProgram->setDepartment(null);
            }
        }

        return $this;
    }

    /**
     * @return Collection<int, Student>
     */
    public function getStudents(): Collection
    {
        return $this->students;
    }

    public function addStudent(Student $student): static
    {
        if (!$this->students->contains($student)) {
            $this->students->add($student);
            $student->setDepartment($this);
        }

        return $this;
    }

    public function removeStudent(Student $student): static
    {
        if ($this->students->removeElement($student)) {
            // set the owning side to null (unless already changed)
            if ($student->getDepartment() === $this) {
                $student->setDepartment(null);
            }
        }

        return $this;
    }
}
