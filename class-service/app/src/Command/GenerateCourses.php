<?php

namespace App\Command;

use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand('app:generate-courses')]
class GenerateCourses extends Command
{
    public function __construct(
        private readonly EntityManagerInterface $entityManager
    )
    {
        parent::__construct();
    }

    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        $conn = $this->entityManager->getConnection();
        $courses = json_decode(file_get_contents(__DIR__ . '/courses.json'),true);

        foreach ($courses as $course) {
            $output->writeln('Inserting ' . $course['name']);
            $sql = <<<SQL
                INSERT INTO course (`id`, `name`, `teacher_id`) VALUES (:id, :name, :teacherId);
            SQL;

            $conn->executeQuery($sql, ['name' => $course['name'], 'teacherId' => rand(1, 10), 'id' => $course['id']]);

            foreach ($course['studyPrograms'] as $studyProgram) {
                $sql = <<<SQL
                    INSERT INTO course_study_program(`course_id`, `study_program_id`) VALUES (:courseId, :studyProgramId);
                SQL;

                $conn->executeQuery($sql, ['courseId' => $course['id'], 'studyProgramId' => $studyProgram]);
            }
        }

        return Command::SUCCESS;
    }
}