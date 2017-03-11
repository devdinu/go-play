#!/bin/sh

function lint {
    errors=$(glide novendor | xargs -I {} $GOBIN/golint {} | grep -v "exported .* should have comment or be unexported")
    echo $errors
    ERROR=`echo $errors | wc -w`
    [ "$ERROR" -eq "0" ] && echo "No Lint Errors!"
    return $ERROR
}

// START OMIT

function coverage {
    PROJECT_DIR=$PWD
    find $PROJECT_DIR -iname 'coverage*.out' -delete
    echo "Generating Coverage Report: $PROJECT_DIR"
	echo "mode: count" > $PROJECT_DIR/coverage-all.out
    for dir in $(go list ./... | grep -v vendor); do (cd $GOSRC/$dir && go test -coverprofile=coverage.out -covermode=count); done
    find . -iname 'coverage.out' -exec grep -v 'mode:' {} \; >> $PROJECT_DIR/coverage-all.out
	cd $PROJECT_DIR && go tool cover -html=$PROJECT_DIR/coverage-all.out
}
// END OMIT

echo "Running golint..."
lint

echo "\nRunning Vet..."
go vet $(glide novendor)

[ "$1" = "-nocover" ] || coverage

echo "Done."
