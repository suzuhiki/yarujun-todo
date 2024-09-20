class Task {
  Task({
    required this.title,
    required this.memo,
    required this.deadline,
    this.waitlist_num = -1,
  });

  final String title;
  final String memo;
  final String deadline;
  final int waitlist_num;

  factory Task.fromJson(dynamic json) {
    return Task(
      title: json['Title'] as String,
      memo: json['Memo'] as String,
      deadline: json['Deadline'] as String,
      waitlist_num: int.parse(json['Waitlist_num'] as String),
    );
  }
}
