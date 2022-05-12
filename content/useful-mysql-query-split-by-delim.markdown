<!--
page_title: Useful MySQL Query | Split by Delim
page_description: A handly little MySQL trick I have just found a use for.
page_status: published
page_date: 2022/05/12
-->

# A Handly Little MySQL Trick

I recently needed to split a string by delimiter,
in this case a full stop `.` character, and found
the following extremely useful to do this:

```

SELECT
  REPLACE (
    LEFT (
        string_column,
        LOCATE (
          '.', string_column
        )
      ),
  '.', '')
FROM your_table

```

This nicely changes:

> Desired Text.--------.Undesired Text

To:

> Desired Text

After subsequently looking it up online, it looks
like there are quite a few functions similar to
`LOCATE` (`SUBSTRING_INDEX`, for example, looks
similar). But `LOCATE` was the first one I
encountered, reading around potentially relevant
functions, so `LOCATE` works for me!

> 12/05/2022:
> Added `blockquote` styling!
