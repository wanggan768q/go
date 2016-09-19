@echo off
Setlocal enabledelayedexpansion
::CODER BY Ambition POWERD BY iBAT

cd %~dp0

cd ..

copy /y .\\UploadFiles\\Skill_技能基础表.xlsx F:\\Work\\config\\Skill_技能基础表.xlsx

start F:\\Work\\cscommon\\build\\export_config.py
