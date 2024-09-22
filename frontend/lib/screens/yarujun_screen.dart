import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:frontend/global.dart';
import 'package:frontend/models/task.dart';
import 'package:frontend/models/api_return.dart';
import 'package:frontend/screens/login_screen.dart';

class YarujunScreen extends StatefulWidget {
  const YarujunScreen({Key? key}) : super(key: key);

  @override
  State<YarujunScreen> createState() => _YarujunScreenState();
}

class _YarujunScreenState extends State<YarujunScreen> {
  List<int> taskIds = [];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('やる順 Todo'),
        automaticallyImplyLeading: false,
      ),
      body: FutureBuilder(
        future: getWaitList(),
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(child: CircularProgressIndicator());
          } else if (snapshot.hasData) {
            var data = snapshot.data as ApiReturn;

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
              final body = data.body as List<Task>;
              taskIds = body.map((e) => int.parse(e.id)).toList();
              return ReorderableListView.builder(
                onReorder: (oldIndex, newIndex) {
                  if (oldIndex < newIndex) {
                    newIndex -= 1;
                  }
                  final int item = taskIds.removeAt(oldIndex);
                  taskIds.insert(newIndex, item);

                  reorderWaitlist().then((value) {
                    if (value.statusCode != 200) {
                      if (value.statusCode == 401) {
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
                      setState(() {});
                    }
                  });
                },
                itemCount: data.body.length,
                itemBuilder: (context, index) {
                  return Container(
                    color: data.body[index].done
                        ? Colors.grey[300]
                        : Colors.transparent,
                    key: Key(data.body[index].id.toString()),
                    child: ListTile(
                      title: Text(data.body[index].title),
                      leading: Text(
                          (data.body[index].waitlistNum + 1).toString(),
                          style: const TextStyle(fontSize: 18)),
                      trailing: const Icon(Icons.drag_handle),
                    ),
                  );
                },
              );
            }
          } else {
            return const Center(child: Text("waitlistを表示できません"));
          }
        },
      ),
    );
  }

  // タスク一覧を取得
  Future<ApiReturn> getWaitList() async {
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
      'filter': "waitlist",
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

  // waitlistを入れ替える
  Future<ApiReturn> reorderWaitlist() async {
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
    };
    final body = jsonEncode(<String, dynamic>{
      'ids': taskIds,
    });
    final url = Uri.parse(
        '$BaseURL/api/v1/auth/tasks/waitlist/reorder?${Uri(queryParameters: query).query}');
    final response = await http.put(url, headers: header, body: body);

    print(url);
    print(response.statusCode);
    print(response.body);

    if (response.statusCode == 200) {
      return ApiReturn(statusCode: 200, body: "Success");
    } else {
      throw Exception('Failed to delete task');
    }
  }
}
