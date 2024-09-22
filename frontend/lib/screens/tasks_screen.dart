import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:frontend/models/api_return.dart';
import 'package:frontend/screens/login_screen.dart';
import 'package:http/http.dart' as http;
import 'package:frontend/global.dart';
import 'package:frontend/models/task.dart';

class TasksScreen extends StatefulWidget {
  const TasksScreen({Key? key}) : super(key: key);

  @override
  State<TasksScreen> createState() => _TasksScreenState();
}

class _TasksScreenState extends State<TasksScreen> {
  DateTime _taskDate = new DateTime.now();
  String _taskSort = "waitlist_num";
  final _formKey = GlobalKey<FormState>();

  String _taskTitle = "";
  int _toggleSelected = 0;

  @override
  Widget build(BuildContext context) {
    final List<bool> selected = [false, false];
    selected[_toggleSelected] = true;

    return Scaffold(
      appBar: AppBar(
        title: const Text("タスク一覧"),
        actions: [
          ToggleButtons(
            isSelected: selected,
            onPressed: (int index) {
              setState(() {
                _toggleSelected = index;
                if (index == 0) {
                  _taskSort = "waitlist_num";
                } else {
                  _taskSort = "deadline";
                }
              });
            },
            borderRadius: const BorderRadius.all(Radius.circular(10)),
            children: const [
              Icon(Icons.format_list_numbered),
              Icon(Icons.calendar_today),
            ],
          ),
          const SizedBox(width: 10),
        ],
        automaticallyImplyLeading: false,
      ),
      body: FutureBuilder(
        future: getTaskList(),
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(child: CircularProgressIndicator());
          } else if (snapshot.hasData) {
            var data = snapshot.data as ApiReturn;
            print(data);

            if (data.statusCode != 200) {
              if (data.statusCode == 401) {
                WidgetsBinding.instance.addPostFrameCallback(
                  (_) {
                    Navigator.pushReplacement(
                      context,
                      MaterialPageRoute(
                          builder: (context) => const LoginScreen()),
                    );
                  },
                );
                return const Center(child: Text("Unauthorized"));
              } else {
                return const Center(child: Text("Error"));
              }
            } else {
              return ListView.builder(
                itemCount: data.body.length,
                itemBuilder: (context, index) {
                  return GestureDetector(
                    child: Card(
                      child: ListTile(
                        title: Text(data.body[index].title),
                        leading: data.body[index].done
                            ? Checkbox(
                                value: true,
                                onChanged: (value) {
                                  putTaskStatus(data.body[index].id, false)
                                      .then(
                                    (value) {
                                      if (value.statusCode != 200) {
                                        if (value.statusCode == 401) {
                                          WidgetsBinding.instance
                                              .addPostFrameCallback(
                                            (_) {
                                              Navigator.pushReplacement(
                                                context,
                                                MaterialPageRoute(
                                                    builder: (context) =>
                                                        const LoginScreen()),
                                              );
                                            },
                                          );
                                          return const Center(
                                              child: Text("Unauthorized"));
                                        } else {
                                          return const Center(
                                              child: Text("Error"));
                                        }
                                      } else {
                                        setState(() {});
                                      }
                                    },
                                  );
                                },
                              )
                            : Checkbox(
                                value: false,
                                onChanged: (value) {
                                  putTaskStatus(data.body[index].id, true).then(
                                    (value) {
                                      if (value.statusCode != 200) {
                                        if (value.statusCode == 401) {
                                          WidgetsBinding.instance
                                              .addPostFrameCallback(
                                            (_) {
                                              Navigator.pushReplacement(
                                                context,
                                                MaterialPageRoute(
                                                    builder: (context) =>
                                                        const LoginScreen()),
                                              );
                                            },
                                          );
                                          return const Center(
                                              child: Text("Unauthorized"));
                                        } else {
                                          return const Center(
                                              child: Text("Error"));
                                        }
                                      } else {
                                        setState(() {});
                                      }
                                    },
                                  );
                                },
                              ),
                        trailing: Text(data.body[index].deadline),
                      ),
                    ),
                    onTap: () {
                      showModalBottomSheet(
                        context: context,
                        builder: (context) {
                          return Container(
                            height: 200,
                            width: MediaQuery.sizeOf(context).width,
                            padding: const EdgeInsets.all(16),
                            child: Column(
                              mainAxisAlignment: MainAxisAlignment.center,
                              mainAxisSize: MainAxisSize.min,
                              children: [
                                Text(data.body[index].title),
                                Text(data.body[index].deadline),
                                Text(data.body[index].waitlistNum.toString()),
                              ],
                            ),
                          );
                        },
                      );
                    },
                  );
                },
              );
            }
          } else {
            print(snapshot.error);
            return const Center(child: Text("タスクを表示できません"));
          }
        },
      ),
      floatingActionButton: FloatingActionButton(
        child: const Icon(Icons.add),
        onPressed: () {
          showModalBottomSheet(
            context: context,
            builder: (context) {
              return Container(
                height: 300 + MediaQuery.of(context).viewInsets.bottom,
                width: MediaQuery.sizeOf(context).width,
                padding: const EdgeInsets.only(top: 2, left: 16, right: 16),
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.start,
                  children: [
                    Form(
                      key: _formKey,
                      child: TextFormField(
                        autofocus: true,
                        validator: (value) {
                          if (value == null || value.isEmpty) {
                            return 'タスク名を入力してください。';
                          }
                          return null;
                        },
                        onSaved: (value) {
                          _taskTitle = value!;
                        },
                        decoration: InputDecoration(
                          hintText: '何をしますか？',
                          suffixIcon: IconButton(
                            onPressed: () {
                              if (_formKey.currentState!.validate()) {
                                _formKey.currentState!.save();
                                if (_taskTitle != "") {
                                  postTask().then(
                                    (value) {
                                      if (value.statusCode != 200) {
                                        if (value.statusCode == 401) {
                                          WidgetsBinding.instance
                                              .addPostFrameCallback(
                                            (_) {
                                              Navigator.pushReplacement(
                                                context,
                                                MaterialPageRoute(
                                                    builder: (context) =>
                                                        const LoginScreen()),
                                              );
                                            },
                                          );
                                          return const Center(
                                              child: Text("Unauthorized"));
                                        } else {
                                          return const Center(
                                              child: Text("Error"));
                                        }
                                      } else {
                                        Navigator.of(context).pop();
                                        setState(() {});
                                      }
                                    },
                                  );
                                } else {
                                  ScaffoldMessenger.of(context).showSnackBar(
                                    const SnackBar(
                                      content: Text('タスク名を入力してください。'),
                                    ),
                                  );
                                }
                              }
                            },
                            icon: const Icon(Icons.send),
                          ),
                        ),
                      ),
                    ),
                    Row(
                      children: [
                        ElevatedButton(
                          child: Row(
                            children: [
                              Icon(Icons.calendar_month),
                              Text(" ${_taskDate.month}/${_taskDate.day}"),
                            ],
                          ),
                          onPressed: () {
                            onPressedRaisedButton();
                          },
                        ),
                        ElevatedButton(
                            onPressed: () {},
                            child: Icon(Icons.format_list_numbered)),
                      ],
                    ),
                  ],
                ),
              );
            },
          );
        },
      ),
    );
  }

  // 日付選択のボタン
  void onPressedRaisedButton() async {
    final picked = await showDatePicker(
            context: context,
            initialDate: _taskDate,
            firstDate: new DateTime(2018),
            lastDate: new DateTime.now().add(new Duration(days: 360)))
        .then((value) {
      if (value != null) {
        setState(() {
          _taskDate = value;
        });
      }
    });
  }

  // タスク一覧を取得
  Future<ApiReturn> getTaskList() async {
    if (Token == "") {
      return ApiReturn(statusCode: 401, body: "Token is empty");
    }

    if (UserID == "") {
      await getUserId().then((value) {
        if (value.statusCode == 200) {
          UserID = value.body;
        } else if (value.statusCode == 401) {
          return ApiReturn(statusCode: 401, body: "Token is empty");
        } else {
          return ApiReturn(statusCode: value.statusCode, body: "Error");
        }
      });
    }

    final query = {
      'user_id': UserID,
      'sort': _taskSort,
    };
    final header = <String, String>{
      'Authorization': 'Bearer $Token',
    };
    final url = Uri.parse(
        '$BaseURL/api/v1/auth/tasks?${Uri(queryParameters: query).query}');
    print(url);
    final response = await http.get(url, headers: header);

    print(response.statusCode);

    if (response.statusCode == 200) {
      final List<dynamic> body = jsonDecode(response.body);
      final List<Task> tasks =
          body.map((dynamic json) => Task.fromJson(json)).toList();
      return ApiReturn(statusCode: 200, body: tasks);
    } else if (response.statusCode == 401) {
      return ApiReturn(statusCode: 401, body: "Unauthorized");
    } else {
      return ApiReturn(statusCode: response.statusCode, body: "Network Error");
    }
  }

  // タスクを追加
  Future<ApiReturn> postTask() async {
    if (Token == "") {
      return ApiReturn(statusCode: 401, body: "Token is empty");
    }

    if (UserID == "") {
      await getUserId().then((value) {
        if (value.statusCode == 200) {
          UserID = value.body;
        } else if (value.statusCode == 401) {
          return ApiReturn(statusCode: 401, body: "Token is empty");
        } else {
          return ApiReturn(statusCode: value.statusCode, body: "Error");
        }
      });
    }

    final query = {
      'user_id': UserID,
    };
    final header = <String, String>{
      'Authorization': 'Bearer $Token',
      'Content-Type': 'application/json',
    };
    final body = jsonEncode(<String, dynamic>{
      'title': _taskTitle,
      'deadline': "${_taskDate.year}-${_taskDate.month}-${_taskDate.day}",
      'waitlist_num': -1,
    });
    print(body);
    final url = Uri.parse(
        '$BaseURL/api/v1/auth/tasks?${Uri(queryParameters: query).query}');
    final response = await http.post(url, body: body, headers: header);

    print(url);
    print(response.statusCode);
    print(response.body);

    if (response.statusCode == 200) {
      return ApiReturn(statusCode: 200, body: "Success");
    } else {
      throw Exception('Failed to create task');
    }
  }

  // タスクのステータスを更新
  Future<ApiReturn> putTaskStatus(String taskId, bool isDone) async {
    if (Token == "") {
      return ApiReturn(statusCode: 401, body: "Token is empty");
    }

    if (taskId == "") {
      return ApiReturn(statusCode: 400, body: "taskId is empty");
    }

    if (UserID == "") {
      await getUserId().then((value) {
        if (value.statusCode == 200) {
          UserID = value.body;
        } else if (value.statusCode == 401) {
          return ApiReturn(statusCode: 401, body: "Token is empty");
        } else {
          return ApiReturn(statusCode: value.statusCode, body: "Error");
        }
      });
    }

    final query = {
      'user_id': UserID,
      'task_id': taskId,
      'status': isDone.toString(),
    };
    final header = <String, String>{
      'Authorization': 'Bearer $Token',
      'Content-Type': 'application/json',
    };
    print(query);
    final url = Uri.parse(
        '$BaseURL/api/v1/auth/tasks/status?${Uri(queryParameters: query).query}');
    final response = await http.put(url, headers: header);

    print(url);
    print(response.statusCode);
    print(response.body);

    if (response.statusCode == 200) {
      return ApiReturn(statusCode: 200, body: "Success");
    } else {
      throw Exception('Failed to create task');
    }
  }
}
