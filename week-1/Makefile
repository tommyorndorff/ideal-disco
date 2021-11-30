all: one two three

one:
	npx @marp-team/marp-cli@latest 01-aws_history.md --pdf --allow-local-files -o pdf/01-aws_history.pdf

two:
	npx @marp-team/marp-cli@latest 02-reliability_pillar.md --pdf --allow-local-files -o pdf/02-reliability_pillar.pdf

three:
	npx @marp-team/marp-cli@latest 03-vpc_and_subnets.md --pdf --allow-local-files -o pdf/03-vpc_and_subnets.pdf

clean:
	rm -f pdf/*.pdf