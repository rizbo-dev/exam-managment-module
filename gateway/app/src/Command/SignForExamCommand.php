<?php

namespace App\Command;

use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;
use Symfony\Contracts\HttpClient\HttpClientInterface;

#[AsCommand(name: "sign-for-exam")]
class SignForExamCommand extends Command
{
    public function __construct(
        private readonly HttpClientInterface $examRegistrationServiceClient
    )
    {
        parent::__construct();
    }

    protected function execute(InputInterface $input, OutputInterface $output): int
    {

        dd($this->examRegistrationServiceClient);
        return Command::SUCCESS;
    }
}