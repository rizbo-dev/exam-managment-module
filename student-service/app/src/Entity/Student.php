<?php

namespace App\Entity;

use ApiPlatform\Metadata\ApiResource;
use App\Repository\StudentRepository;
use Doctrine\DBAL\Types\Types;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: StudentRepository::class)]
#[ApiResource]
class Student
{
    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column]
    private ?int $id = null;

    #[ORM\Column(length: 255)]
    private ?string $firstname = null;

    #[ORM\Column(length: 255)]
    private ?string $lastname = null;

    #[ORM\Column(length: 255)]
    private ?string $jmbg = null;

    #[ORM\Column(length: 255)]
    private ?string $indexNumber = null;

    #[ORM\Column(type: Types::DATE_MUTABLE)]
    private ?\DateTimeInterface $dateOfBirth = null;

    #[ORM\Column]
    private ?int $enrollmentYear = null;

    #[ORM\Column]
    private ?int $currentClassYear = null;

    #[ORM\Column(length: 255)]
    private ?string $address = null;

    #[ORM\Column(length: 255)]
    private ?string $phoneNumber = null;

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getFirstname(): ?string
    {
        return $this->firstname;
    }

    public function setFirstname(string $firstname): static
    {
        $this->firstname = $firstname;

        return $this;
    }

    public function getLastname(): ?string
    {
        return $this->lastname;
    }

    public function setLastname(string $lastname): static
    {
        $this->lastname = $lastname;

        return $this;
    }

    public function getJmbg(): ?string
    {
        return $this->jmbg;
    }

    public function setJmbg(string $jmbg): static
    {
        $this->jmbg = $jmbg;

        return $this;
    }

    public function getIndexNumber(): ?string
    {
        return $this->indexNumber;
    }

    public function setIndexNumber(string $indexNumber): static
    {
        $this->indexNumber = $indexNumber;

        return $this;
    }

    public function getDateOfBirth(): ?\DateTimeInterface
    {
        return $this->dateOfBirth;
    }

    public function setDateOfBirth(\DateTimeInterface $dateOfBirth): static
    {
        $this->dateOfBirth = $dateOfBirth;

        return $this;
    }

    public function getEnrollmentYear(): ?int
    {
        return $this->enrollmentYear;
    }

    public function setEnrollmentYear(int $enrollmentYear): static
    {
        $this->enrollmentYear = $enrollmentYear;

        return $this;
    }

    public function getCurrentClassYear(): ?int
    {
        return $this->currentClassYear;
    }

    public function setCurrentClassYear(int $currentClassYear): static
    {
        $this->currentClassYear = $currentClassYear;

        return $this;
    }

    public function getAddress(): ?string
    {
        return $this->address;
    }

    public function setAddress(string $address): static
    {
        $this->address = $address;

        return $this;
    }

    public function getPhoneNumber(): ?string
    {
        return $this->phoneNumber;
    }

    public function setPhoneNumber(string $phoneNumber): static
    {
        $this->phoneNumber = $phoneNumber;

        return $this;
    }
}
