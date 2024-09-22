class Task {
  Task({
    required this.id,
    required this.title,
    required this.deadline,
    required this.waitlistNum,
    required this.done,
  });

  final String id;
  final String title;
  final String deadline;
  final int waitlistNum;
  final bool done;

  factory Task.fromJson(dynamic json) {
    print(json);
    if (json["Waitlist_num"] == "") {
      return Task(
        id: json['ID'] as String,
        title: json['Title'] as String,
        deadline: json['Deadline'] as String,
        waitlistNum: -1,
        done: json['Done'] as bool,
      );
    } else {
      return Task(
        id: json['Id'] as String,
        title: json['Title'] as String,
        deadline: json['Deadline'] as String,
        waitlistNum: int.parse(json['Waitlist_num'] as String),
        done: json['Done'] as bool,
      );
    }
  }
}
