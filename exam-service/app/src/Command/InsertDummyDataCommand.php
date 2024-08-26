<?php

namespace App\Command;

use App\Entity\Exam;
use App\Entity\ExaminationPeriod;
use DateTime;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand('app:insert-dummy-data')]
class InsertDummyDataCommand extends Command
{
    public function __construct(
        private readonly EntityManagerInterface $entityManager
    )
    {
        parent::__construct();
    }

    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        $data = json_decode(file_get_contents(__DIR__. '/dummyData.json'), true);

        foreach ($data as $item) {
            $examinationPeriod = new ExaminationPeriod();
            $examinationPeriod->setName($item['name'])
                              ->setStartDate(new DateTime($item['start_date']))
                              ->setEndDate(new DateTime($item['end_date']));

            $this->entityManager->persist($examinationPeriod);

            foreach ($item['exams'] as $examItem) {
                $exam = new Exam();
                $exam->setCourseId($examItem['courseId'])
                    ->setMaxStudentEntry($examItem['maxStudentEntry'])
                    ->setCost($examItem['cost']);
                $this->entityManager->persist($exam);

                $examinationPeriod->addExam($exam);
            }
        }

        $this->entityManager->flush();

        return Command::SUCCESS;
    }
}