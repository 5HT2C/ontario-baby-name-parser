# obnp

What? A parser for Ontario's baby name data

Why? I wanted to see if a specific name existed in both the male and female datasets. This tool is more useful than just that.

Where? You can find the [male](https://data.ontario.ca/dataset/ontario-top-baby-names-male) and [female](https://data.ontario.ca/dataset/ontario-top-baby-names-female) data on the Ontario website.

Who? Well, this data includes any name that is registered more than 5 times in one year. If there are less than 6 people with that name registered in a year, they are not included in the data.

## Usage

```bash
git clone git@github.com:l1ving/ontario-baby-name-parser.git
cd ontario-baby-name-parser
make
./obnp -h
```
```bash
Usage of ./obnp:
    -last int
        Limit to last X years (default 0)
    -name string
        Name to find
```