<style>
span.warn {
  color: red;
  font-weight: bold;
}
</style>
<script>

//debugger

function is_digit(character) {
  return character === '0' || character === '1'
      || character === '2' || character === '3'
      || character === '4' || character === '5'
      || character === '6' || character === '7'
      || character === '8' || character === '9';
}

function month_name_(month_number_text) {
  switch (month_number_text) {
  case '01': return 'Jan';
  case '02': return 'Feb';
  case '03': return 'Mar';
  case '04': return 'Apr';
  case '05': return 'May';
  case '06': return 'Jun';
  case '07': return 'Jul';
  case '08': return 'Aug';
  case '09': return 'Sep';
  case '10': return 'Oct';
  case '11': return 'Nov';
  case '12': return 'Dec';
  }
  return -1;
}

/*
function pad(date_day_) {
  if (date_day_.length === 1) {
    return '0' + date_day_;
  }
  return date_day_;
}
*/

function highlight_date(date_text_, cell) {
  var now = new Date();
  var date_year_ = date_text_.substr(0, 4);
  var date_month_ = date_text_.substr(5, 2);
  var date_day_ = date_text_.substr(8, 2);
  var date_month_text_ = month_name_(date_month_);
  var date_day_padded_ = /*pad(*/date_day_/*)*/;
  var date_ = Date.parse(date_day_padded_ + ' ' + date_month_text_ + ' ' + date_year_ + ' 06:00:00 GMT');
  if (date_ < now) {
    cell.innerHTML = '<span class="warn">' + cell.innerHTML + '</span>';
  }
}

function highlight_possible_date(maybe_date_, cell) {
  //alert(maybe_date_);
  if (maybe_date_.length === 10) {
    if (is_digit(maybe_date_[0]) && is_digit(maybe_date_[1])
        && is_digit(maybe_date_[2]) && is_digit(maybe_date_[3])
        && maybe_date_[4] === '-'
        && is_digit(maybe_date_[5]) && is_digit(maybe_date_[6])
        && maybe_date_[7] === '-'
        && is_digit(maybe_date_[8]) && is_digit(maybe_date_[9])) {
      highlight_date(maybe_date_, cell);
    }
  }
}

function highlight_dates(table) {
  //alert(table.lastElementChild);
  var tbody = table.lastElementChild;
  var rows = tbody.children;
  var row;
  for (var i = 0; i < rows.length; i++) {
    row = rows[i];
    var cells = row.children;
    var cell;
    for (var j = 0; j < cells.length; j++) {
      if (j > 1) {
        var cell = cells[j];
        var maybe_date_ = cell.innerHTML;
        highlight_possible_date(maybe_date_, cell);
      }
    }
  }
}

function work_table(table) {
  if (table.firstElementChild) {
    var thead = table.firstElementChild;
    if (thead.firstElementChild) {
      var theadr = thead.firstElementChild;
      if (theadr.children && theadr.children.length === 4) {
        if (theadr.children[0].innerHTML === 'Version'
            && theadr.children[1].innerHTML === 'Release Date'
            && theadr.children[2].innerHTML === 'Maintenance End'
            && theadr.children[3].innerHTML === 'End of Life') {
          highlight_dates(table);
        }
      }
    }
  }
}

function work_tables(tables) {
  var i;
  for (i = 0; i < tables.length; i++) {
    work_table(tables[i]);
  }
}

var tables = document.getElementsByTagName('table');

if (tables.length > 0) {
  work_tables(tables);
}

</script>
