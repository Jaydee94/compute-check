# Fundamental architecture

This document describes the planned architecture of this application for initial development.

## Components

This application consists of multiple componentes.

### Calculator

The calculating component is used for calculating optimal resource quota for a kubernetes namespace. 

#### Functions
|Description|Core-Function|Addon|
|-----------|-------------|-----|
|calculate limit cpu|x| |
|calculate limit memory |x||
|calculate request cpu|x||
|calculate request memory|x||
|calculate request gpu||x|

### Analyzer

The analyzing component analyzes the configured resource quota for namespaces and compares it with the current state of pods running in this namespace.

#### Functions

|Description|Core-Function|Addon|
|-----------|-------------|-----|
|analyze running pod per namespace|x||
|analyze configured resource quota per namespace|x||
|provide a rendered resource quota yaml file with optimized values||x|
|provide the diff of optimal quota compared to the current state|x||

### UI

The ui provides the frontend application for using the componentes `calculator` and `analyzer`.

#### Functions

|Description|Core-Function|Addon|
|-----------|-------------|-----|
|index tabs for switching between calculator and analyzer|x||
|dynamic form to provide the planned object of a kubernetes namespace|x||
|analyze complete cluster|x||
|analyze single namespace|x||
|display rendered compute resources yaml file with optimized values||x|