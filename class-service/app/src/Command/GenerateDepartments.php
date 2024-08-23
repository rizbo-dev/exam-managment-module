<?php

namespace App\Command;

use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand('app:generate-departments')]
class GenerateDepartments extends Command
{
    public function __construct(
        private readonly EntityManagerInterface $entityManager
    )
    {
        parent::__construct();
    }

    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        $departments = json_decode(file_get_contents(__DIR__ . '/departments.json'), true);
        $conn = $this->entityManager->getConnection();

        foreach ($departments as $department)
        {
            $output->writeln("Inserting: " . $department['name']);
            $sql = <<<SQL
                INSERT INTO department (`id`, `name`) VALUES (:id, :value);
            SQL;
            $conn->executeQuery($sql, ['id' => $department['id'], 'value' => $department['name']]);

            foreach ($department['studyPrograms'] as $studyProgram) {
                $sql = <<<SQL
                    INSERT INTO study_program (`name`, `department_id`) VALUES (:value, :department_id);
                SQL;

                $conn->executeQuery($sql, ['department_id' => $department['id'], 'value' => $studyProgram]);
            }
        }


        return Command::SUCCESS;
    }
}