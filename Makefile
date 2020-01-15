GO=GO
SWIG=SWIG
SWIGFLAGS=-go -c++ -intgosize 64

all: clean install

clean:
	rm -rf goparquet_wrap.cxx
	rm -rf goparquet_wrap.h
	rm -rf goparquet.go
	rm -rf goparquet.h
	rm -rf goparquet.cxx

install: clean
	$(SWIG) $(SWIGFLAGS) ./swig/goparquet.swigcxx
	mv ./swig/goparquet_wrap.cxx .
	mv ./swig/goparquet_wrap.h .
	mv ./swig/goparquet.go .
	cp ./swig/goparquet.h .
	cp ./swig/goparquet.cxx .
	$(GO) install
