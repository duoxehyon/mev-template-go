import sys


def main(command):
    if command == "test":
        # do testing
        print("testing")
    elif command == "build":
        print("building")
        import os
        os.system('go build ./cmd/bot/main.go')
if len(sys.argv) == 1:
    print("test / build")
    exit(1)
main(sys.argv[1])