#Set this to @ to keep the makefile quiet
ifndef SILENCE
	SILENCE = @
endif

#--- Inputs ----#
COMPONENT_NAME = CodingExercises1
CPPUTEST_HOME = cpputest
CPPUTEST_WARNINGFLAGS =  -Wall

CPPUTEST_USE_EXTENSIONS = Y
CPP_PLATFORM = Gcc

SRC_DIRS = \
	src

TEST_SRC_DIRS = \
	tst

INCLUDE_DIRS =\
  .\
  src\
  $(CPPUTEST_HOME)/include\

include $(CPPUTEST_HOME)/build/MakefileWorker.mk


