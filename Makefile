include $(GOROOT)/src/Make.inc

TARG=linear

GOFILES=linear.go

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
