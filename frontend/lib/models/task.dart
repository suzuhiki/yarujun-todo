class Task {
  Task({
    required this.id,
    required this.title,
    required this.deadline,
    this.waitlistNum = -1,
  });

  final String id;
  final String title;
  final String deadline;
  final int waitlistNum;

  factory Task.fromJson(dynamic json) {
    if (json["Waitlist_num"] == "") {
      return Task(
        id: json['ID'] as String,
        title: json['Title'] as String,
        deadline: json['Deadline'] as String,
        waitlistNum: -1,
      );
    } else {
      return Task(
        id: json['ID'] as String,
        title: json['Title'] as String,
        deadline: json['Deadline'] as String,
        waitlistNum: int.parse(json['Waitlist_num'] as String),
      );
    }
  }
}
