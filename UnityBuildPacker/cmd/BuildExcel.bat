@echo off
Setlocal enabledelayedexpansion
::CODER BY Ambition POWERD BY iBAT

cd %~dp0

cd ..

copy /y .\\UploadFiles\\Skill_���ܻ�����.xlsx F:\\Work\\config\\Skill_���ܻ�����.xlsx

start F:\\Work\\cscommon\\build\\export_config.py
